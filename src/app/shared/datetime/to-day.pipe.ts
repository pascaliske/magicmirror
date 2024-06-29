import { Pipe, PipeTransform } from '@angular/core'
import { getDate } from 'date-fns'

@Pipe({
    standalone: true,
    name: 'toDay',
})
export class ToDayPipe implements PipeTransform {
    public transform(value: Date): number {
        return getDate(value)
    }
}
