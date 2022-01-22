import { Component, ChangeDetectionStrategy } from '@angular/core'
import { forkJoin, Observable, OperatorFunction } from 'rxjs'
import { concatMap, map } from 'rxjs/operators'
import { Cacheable } from 'ts-cacheable'
import { BaseCardComponent } from '../base-card/base-card.component'

@Component({
    selector: 'cmp-news-card',
    templateUrl: './news-card.component.html',
    styleUrls: ['./news-card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NewsCardComponent extends BaseCardComponent {
    public readonly data$: Observable<string[]> = this.settings$.pipe(
        concatMap(({ data }) => this.fetchNewsData(data?.feeds ?? [])),
    )

    @Cacheable({ maxAge: 600000 })
    private fetchNewsData(urls: string[]): Observable<string[]> {
        const feeds: Observable<string[]>[] = urls.map((url: string) => {
            return this.http.get(url, { responseType: 'text' }).pipe(this.extractHeadlines(4))
        })

        return forkJoin(feeds).pipe(map(responses => responses.flat()))
    }

    private extractHeadlines(count: number): OperatorFunction<string, string[]> {
        return map<string, string[]>((response: string) => {
            const parser: DOMParser = new window.DOMParser()
            const feed: Document = parser.parseFromString(response, 'text/xml')

            return Array.from(feed.querySelectorAll('item > title'))
                .map(item => item?.innerHTML ?? '')
                .slice(0, count)
        })
    }
}
