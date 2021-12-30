import { Component, Input, ChangeDetectionStrategy } from '@angular/core'
import { IconDefinition } from '@fortawesome/free-solid-svg-icons'

@Component({
    selector: 'cmp-icon',
    templateUrl: './icon.component.html',
    styleUrls: ['./icon.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IconComponent {
    @Input()
    public icon!: IconDefinition

    @Input()
    public size!: 'xs' | 'sm' | 'lg' | '2x' | '3x' | '5x' | '7x' | '10x'

    @Input()
    public animate: 'spin' | 'pulse' | undefined
}
