import { Pipe, PipeTransform, Inject, LOCALE_ID } from '@angular/core'

@Pipe({
    name: 'toLocalDate',
})
export class ToLocalDatePipe implements PipeTransform {
    public constructor(@Inject(LOCALE_ID) private readonly locale: string) {}

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
