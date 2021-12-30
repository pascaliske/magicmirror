import { Pipe, PipeTransform } from '@angular/core'
import { getMonth } from 'date-fns'

@Pipe({
    name: 'toMonth',
})
export class ToMonthPipe implements PipeTransform {
    public transform(value: Date): number {
        return getMonth(value)
    }
}
