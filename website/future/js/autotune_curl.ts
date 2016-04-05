import 'rxjs/Rx';
import {Component} from 'angular2/core';
import {NgForm}    from 'angular2/common';
import {HTTP_PROVIDERS, Http} from 'angular2/http';

@Component({
    selector: 'autotune-curl',
    viewProviders: [HTTP_PROVIDERS],
    // Location of the template for this component
    template: `
<div>
<label>Server Profile:</label>
<select [(ngModel)]="selectedProfile" class="form-control">
<option *ngFor="#p of profiles" [value]="p">{{p}}</option>
</select>
</div>
<br>
<code>\\curl -sSL https://acksin.com/autotune/install.sh | bash -s {{selectedProfile}}</code>`
})
export class AutotuneCurl {
    selectedProfile: string = 'apache';

    constructor(http: Http) {
        http.get('/autotune/js/profiles.json')
            .map(res => res.json())
            .subscribe(profiles => this.profiles = profiles);
    }
}
