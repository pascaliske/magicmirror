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
        concatMap(({ data }) => {
            const feeds = data?.feeds?.map(url => {
                return this.fetchNewsData(url).pipe(this.extractHeadlines(4))
            })

            return forkJoin(feeds).pipe(map(responses => responses.flat()))
        }),
    )

    @Cacheable({ maxAge: 600000 })
    private fetchNewsData(url: string): Observable<string> {
        return this.http.get(url, { responseType: 'text' })
    }

    private extractHeadlines(count: number): OperatorFunction<string, string[]> {
        return map<string, string[]>(response => {
            const parser: DOMParser = new window.DOMParser()
            const feed: Document = parser.parseFromString(response, 'text/xml')

            return Array.from(feed.querySelectorAll('item > title'))
                .map(item => item?.innerHTML ?? '')
                .slice(0, count)
        })
    }
}
