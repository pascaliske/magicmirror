import { Pipe, PipeTransform, LOCALE_ID, inject } from '@angular/core'

@Pipe({
    standalone: true,
    name: 'toLocalDate',
})
export class ToLocalDatePipe implements PipeTransform {
    private readonly locale: string = inject(LOCALE_ID)

    public transform(value: Date, length: 'short' | 'long' = 'short'): string {
        if (length === 'short') {
            return value.toLocaleDateString(this.locale.toLowerCase(), {
                day: '2-digit',
                month: '2-digit',
                year: 'numeric',
            })
        }

        return value.toLocaleDateString(this.locale.toLowerCase(), {
            weekday: 'long',
            day: 'numeric',
            month: 'long',
            year: 'numeric',
        })
    }
}
