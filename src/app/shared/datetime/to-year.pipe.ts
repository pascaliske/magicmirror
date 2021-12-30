import { Pipe, PipeTransform } from '@angular/core'
import { getYear } from 'date-fns'

@Pipe({
    name: 'toYear',
})
export class ToYearPipe implements PipeTransform {
    public transform(value: Date): number {
        return getYear(value)
    }
}
