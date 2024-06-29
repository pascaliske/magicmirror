import { Component, ChangeDetectionStrategy } from '@angular/core'
import { NgIf, AsyncPipe } from '@angular/common'
import { TranslateModule } from '@ngx-translate/core'
import { Observable, interval } from 'rxjs'
import { startWith, map } from 'rxjs/operators'
import { ToWeekPipe } from 'shared/datetime/to-week.pipe'
import { ToLocalDatePipe } from 'shared/datetime/to-local-date.pipe'
import { ToLocalTimePipe } from 'shared/datetime/to-local-time.pipe'
import { BaseCardComponent } from '../base-card/base-card.component'

@Component({
    standalone: true,
    selector: 'cmp-time-card',
    templateUrl: './time-card.component.html',
    styleUrls: ['./time-card.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [NgIf, AsyncPipe, TranslateModule, ToWeekPipe, ToLocalDatePipe, ToLocalTimePipe],
})
export class TimeCardComponent extends BaseCardComponent {
    public readonly interval$: Observable<Date> = interval(1000).pipe(
        startWith(new Date()),
        map(() => new Date()),
    )
}
