import { Component, Input, ChangeDetectionStrategy } from '@angular/core'
import { NgIf } from '@angular/common'
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome'
import { IconProp } from '@fortawesome/fontawesome-svg-core'

@Component({
    standalone: true,
    selector: 'cmp-icon',
    templateUrl: './icon.component.html',
    styleUrls: ['./icon.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [NgIf, FontAwesomeModule],
})
export class IconComponent {
    @Input()
    public icon!: IconProp

    @Input()
    public size!: 'xs' | 'sm' | 'lg' | '2x' | '3x' | '5x' | '7x' | '10x'

    @Input()
    public animate: 'spin' | 'pulse' | undefined
}
