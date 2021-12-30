import { Pipe, PipeTransform, Inject } from '@angular/core'
import { format } from 'date-fns'
import { DateFnsConfig, DATE_FNS_CONFIG } from './config'

@Pipe({
    name: 'toMonthName',
})
export class ToMonthNamePipe implements PipeTransform {
    public constructor(@Inject(DATE_FNS_CONFIG) private readonly dateFnsConfig: DateFnsConfig) {}

    public transform(value: Date): string {
        return format(value, 'LLLL', this.dateFnsConfig)
    }
}
