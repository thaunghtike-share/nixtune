import {Component} from 'angular2/core';
import {NgForm}    from 'angular2/common';

@Component({
    // Declare the tag name in index.html to where the component attaches
    selector: 'autotune-curl',
    // Location of the template for this component
    template: `
<div>
<label>Server Profile:</label>
<select [(ngModel)]="selectedProfile" class="form-control">
<option *ngFor="#p of profiles" [value]="p">{{p}}</option>
</select>
</div>
<br>
<code>\\curl -sSL https://anatma.co/autotune/install.sh | bash -s {{selectedProfile}}</code>`
})
export class AutotuneCurl {
    profiles: []string = [
        "golang",
        "nodejs",
        "haproxy"
    ];
    // Default profile to use.
    selectedProfile: string = 'golang';
}
