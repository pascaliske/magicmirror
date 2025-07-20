import { Pipe, PipeTransform, inject } from '@angular/core'
import { format } from 'date-fns'
import { DateFnsConfig, DATE_FNS_CONFIG } from './config'

@Pipe({
    standalone: true,
    name: 'toMonthName',
})
export class ToMonthNamePipe implements PipeTransform {
    private readonly dateFnsConfig: DateFnsConfig = inject<DateFnsConfig>(DATE_FNS_CONFIG)

    public transform(value: Date): string {
        return format(value, 'LLLL', this.dateFnsConfig)
    }
}
