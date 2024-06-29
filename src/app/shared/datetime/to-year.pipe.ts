import { Pipe, PipeTransform } from '@angular/core'
import { getYear } from 'date-fns'

@Pipe({
    standalone: true,
    name: 'toYear',
})
export class ToYearPipe implements PipeTransform {
    public transform(value: Date): number {
        return getYear(value)
    }
}
