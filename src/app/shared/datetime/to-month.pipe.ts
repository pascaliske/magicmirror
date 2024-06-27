import { Pipe, PipeTransform } from '@angular/core'
import { getMonth } from 'date-fns'

@Pipe({
    standalone: true,
    name: 'toMonth',
})
export class ToMonthPipe implements PipeTransform {
    public transform(value: Date): number {
        return getMonth(value)
    }
}
