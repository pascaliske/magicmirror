import { NgModule } from '@angular/core'
import { CommonModule } from '@angular/common'
import { TranslateModule } from '@ngx-translate/core'
import { FromTimestampPipe } from './datetime/from-timestamp.pipe'
import { ToLocalDatePipe } from './datetime/to-local-date.pipe'
import { ToLocalTimePipe } from './datetime/to-local-time.pipe'
import { ToDayPipe } from './datetime/to-day.pipe'
import { ToDayNamePipe } from './datetime/to-day-name.pipe'
import { ToWeekPipe } from './datetime/to-week.pipe'
import { ToMonthPipe } from './datetime/to-month.pipe'
import { ToMonthNamePipe } from './datetime/to-month-name.pipe'
import { ToYearPipe } from './datetime/to-year.pipe'
import { ToHumanNumberPipe } from './numbers/to-human.pipe'

@NgModule({
    imports: [CommonModule, TranslateModule],
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
        ToHumanNumberPipe,
    ],
    exports: [
        CommonModule,
        TranslateModule,
        FromTimestampPipe,
        ToLocalDatePipe,
        ToLocalTimePipe,
        ToDayPipe,
        ToDayNamePipe,
        ToWeekPipe,
        ToMonthPipe,
        ToMonthNamePipe,
        ToYearPipe,
        ToHumanNumberPipe,
    ],
})
export class SharedModule {}
