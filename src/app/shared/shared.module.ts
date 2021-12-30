import { NgModule } from '@angular/core'
import { CommonModule } from '@angular/common'
import { FromTimestampPipe } from './datetime/from-timestamp.pipe'
import { ToLocalDatePipe } from './datetime/to-local-date.pipe'
import { ToLocalTimePipe } from './datetime/to-local-time.pipe'
import { ToDayPipe } from './datetime/to-day.pipe'
import { ToDayNamePipe } from './datetime/to-day-name.pipe'
import { ToWeekPipe } from './datetime/to-week.pipe'
import { ToMonthPipe } from './datetime/to-month.pipe'
import { ToMonthNamePipe } from './datetime/to-month-name.pipe'
import { ToYearPipe } from './datetime/to-year.pipe'

@NgModule({
    imports: [CommonModule],
    declarations: [
        FromTimestampPipe,
        ToLocalDatePipe,
        ToLocalTimePipe,
        ToDayPipe,
        ToDayNamePipe,
        ToWeekPipe,
        ToMonthPipe,
        ToMonthNamePipe,
        ToYearPipe,
    ],
    exports: [
        CommonModule,
        FromTimestampPipe,
        ToLocalDatePipe,
        ToLocalTimePipe,
        ToDayPipe,
        ToDayNamePipe,
        ToWeekPipe,
        ToMonthPipe,
        ToMonthNamePipe,
        ToYearPipe,
    ],
})
export class SharedModule {}
