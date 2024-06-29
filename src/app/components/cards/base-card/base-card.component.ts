import { Component, ChangeDetectionStrategy, Inject, LOCALE_ID } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Store, select } from '@ngrx/store'
import { Observable, OperatorFunction, interval } from 'rxjs'
import { filter, concatMap, startWith, map } from 'rxjs/operators'
import { SettingsFeature, SettingsState } from 'store/settings'

@Component({
    standalone: true,
    selector: 'cmp-base-card',
    templateUrl: './base-card.component.html',
    styleUrls: ['./base-card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class BaseCardComponent {
    protected settings$: Observable<SettingsState> = this.store.pipe(
        select(SettingsFeature.selectSettingsState),
        filter(({ loaded }) => loaded),
    )

    public constructor(
        @Inject(LOCALE_ID) protected readonly locale: string,
        protected http: HttpClient,
        private readonly store: Store,
    ) {}
}

/**
 * Repeats the source observable periodically after the given amount of time.
 *
 * @param time Interval between subsequent repeats in ms. Defaults to 10 minutes.
 * @returns Returns an {@link OperatorFunction} usable in observable pipes.
 */
export function repeatAfter<T>(time: number = 10 * 60 * 1000): OperatorFunction<T, T> {
    return concatMap<T, Observable<T>>((data: T) => {
        return interval(time).pipe(
            startWith(data),
            map(() => data),
        )
    })
}
