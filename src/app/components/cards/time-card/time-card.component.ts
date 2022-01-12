import { Component, ChangeDetectionStrategy } from '@angular/core'
import { Observable, interval } from 'rxjs'
import { startWith, map } from 'rxjs/operators'
import { BaseCardComponent } from '../base-card/base-card.component'

@Component({
    selector: 'cmp-time-card',
    templateUrl: './time-card.component.html',
    styleUrls: ['./time-card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TimeCardComponent extends BaseCardComponent {
    public readonly interval$: Observable<Date> = interval(1000).pipe(
        startWith(new Date()),
        map(() => new Date()),
    )
}
