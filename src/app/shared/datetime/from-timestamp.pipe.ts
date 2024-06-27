import { Pipe, PipeTransform } from '@angular/core'
import { fromUnixTime } from 'date-fns'

@Pipe({
    standalone: true,
    name: 'fromTimestamp',
})
export class FromTimestampPipe implements PipeTransform {
    public transform(value: number): Date {
        return fromUnixTime(value)
    }
}
