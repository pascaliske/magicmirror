import { Pipe, PipeTransform, LOCALE_ID, inject } from '@angular/core'

@Pipe({
    standalone: true,
    name: 'toHumanNumber',
})
export class ToHumanNumberPipe implements PipeTransform {
    private readonly locale: string = inject(LOCALE_ID)

    public transform(value: number, format: string = '1.0-1'): string {
        const [minIntegerDigits, minFractionDigits, maxFractionDigits] = format.split(/[.-]/)

        return value.toLocaleString(this.locale.toLowerCase(), {
            minimumIntegerDigits: parseInt(minIntegerDigits ?? '1', 10),
            minimumFractionDigits: parseInt(minFractionDigits ?? '0', 10),
            maximumFractionDigits: parseInt(maxFractionDigits ?? '1', 10),
        })
    }
}
