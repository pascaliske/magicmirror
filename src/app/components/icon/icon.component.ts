import { Component, ChangeDetectionStrategy, input } from '@angular/core'
import { FaIconComponent } from '@fortawesome/angular-fontawesome'
import { IconProp } from '@fortawesome/fontawesome-svg-core'

@Component({
    standalone: true,
    selector: 'cmp-icon',
    templateUrl: './icon.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [FaIconComponent],
})
export class IconComponent {
    public icon = input<IconProp>()

    public size = input<'xs' | 'sm' | 'lg' | '2x' | '3x' | '5x' | '7x' | '10x'>()

    public animate = input<'beat' | 'fade' | 'beat-fade' | 'bounce' | 'flip' | 'shake' | 'spin'>()
}
