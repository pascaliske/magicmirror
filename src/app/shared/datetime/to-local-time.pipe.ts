import { Pipe, PipeTransform, Inject, LOCALE_ID } from '@angular/core'

@Pipe({
    standalone: true,
    name: 'toLocalTime',
})
export class ToLocalTimePipe implements PipeTransform {
    public constructor(@Inject(LOCALE_ID) private readonly locale: string) {}

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
