import { Pipe, PipeTransform, Inject } from '@angular/core'
import { format } from 'date-fns'
import { DateFnsConfig, DATE_FNS_CONFIG } from './config'

@Pipe({
    standalone: true,
    name: 'toDayName',
})
export class ToDayNamePipe implements PipeTransform {
    public constructor(@Inject(DATE_FNS_CONFIG) private readonly dateFnsConfig: DateFnsConfig) {}

    public transform(value: Date): string {
        return format(value, 'iiii', this.dateFnsConfig)
    }
}
