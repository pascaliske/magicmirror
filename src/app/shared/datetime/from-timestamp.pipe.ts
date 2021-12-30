import { Pipe, PipeTransform } from '@angular/core'
import { fromUnixTime } from 'date-fns'

@Pipe({
    name: 'fromTimestamp',
})
export class FromTimestampPipe implements PipeTransform {
    public transform(value: number): Date {
        return fromUnixTime(value)
    }
}
