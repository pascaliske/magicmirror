import { Pipe, PipeTransform } from '@angular/core'
import { getISOWeek } from 'date-fns'

@Pipe({
    standalone: true,
    name: 'toWeek',
})
export class ToWeekPipe implements PipeTransform {
    public transform(value: Date): string {
        return getISOWeek(value).toString()
    }
}
