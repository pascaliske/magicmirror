import { Component, ChangeDetectionStrategy } from '@angular/core'
import { forkJoin, Observable, OperatorFunction } from 'rxjs'
import { concatMap, map } from 'rxjs/operators'
import { Cacheable } from 'ts-cacheable'
import { BaseCardComponent, repeatAfter } from '../base-card/base-card.component'

interface Headline {
    title: string
    link: string
}

@Component({
    selector: 'cmp-news-card',
    templateUrl: './news-card.component.html',
    styleUrls: ['./news-card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NewsCardComponent extends BaseCardComponent {
    public readonly data$: Observable<Headline[]> = this.settings$.pipe(
        repeatAfter(),
        concatMap(({ data }) => this.fetchNewsData(data?.feeds ?? [])),
    )

    @Cacheable()
    private fetchNewsData(urls: string[]): Observable<Headline[]> {
        const feeds: Observable<Headline[]>[] = urls.map((url: string) => {
            return this.http.get(url, { responseType: 'text' }).pipe(this.extractHeadlines(4))
        })

        return forkJoin(feeds).pipe(map(responses => responses.flat()))
    }

    private extractHeadlines(count: number): OperatorFunction<string, Headline[]> {
        return map<string, Headline[]>((response: string) => {
            const parser: DOMParser = new window.DOMParser()
            const feed: Document = parser.parseFromString(response, 'text/xml')

            return Array.from(feed.querySelectorAll('item'))
                .slice(0, count)
                .map(item => {
                    const title = item?.querySelector('title')?.innerHTML ?? ''
                    const link = item?.querySelector('link')?.innerHTML ?? ''

                    return { title, link }
                })
        })
    }
}
