import { Pipe, PipeTransform, Inject, LOCALE_ID } from '@angular/core'

@Pipe({
    name: 'toHumanNumber',
})
export class ToHumanNumberPipe implements PipeTransform {
    public constructor(@Inject(LOCALE_ID) private readonly locale: string) {}

    public transform(value: number, format: string = '1.0-1'): string {
        const [minIntegerDigits, minFractionDigits, maxFractionDigits] = format.split(/[.-]/)

        return value.toLocaleString(this.locale.toLowerCase(), {
            minimumIntegerDigits: parseInt(minIntegerDigits ?? '1', 10),
            minimumFractionDigits: parseInt(minFractionDigits ?? '0', 10),
            maximumFractionDigits: parseInt(maxFractionDigits ?? '1', 10),
        })
    }
}
