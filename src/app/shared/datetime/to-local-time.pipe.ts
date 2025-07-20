import { Pipe, PipeTransform, LOCALE_ID, inject } from '@angular/core'

@Pipe({
    standalone: true,
    name: 'toLocalTime',
})
export class ToLocalTimePipe implements PipeTransform {
    private readonly locale: string = inject(LOCALE_ID)

    public transform(value: Date, length: 'short' | 'long' = 'short'): string {
        if (length === 'short') {
            return value.toLocaleTimeString(this.locale.toLowerCase(), {
                hour: 'numeric',
                minute: '2-digit',
                hour12: this.locale.toLowerCase() === 'en',
            })
        }

        return value.toLocaleTimeString(this.locale.toLowerCase(), {
            hour: 'numeric',
            minute: '2-digit',
            second: '2-digit',
            hour12: this.locale.toLowerCase() === 'en',
        })
    }
}
