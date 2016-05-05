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
<label>Server Signature:</label>
<select [(ngModel)]="selectedSignature" #signature (change)="selected($event, signature.value)" class="form-control">
<option *ngFor="#p of signatures" [value]="p">{{p}}</option>
</select>
</div>
<br>
<code>\\curl -sSL https://acksin.com/autotune/install.sh | bash -s {{selectedSignature}}</code>

{{premiumSig}}
`
})
export class AutotuneCurl {
    selectedSignature: string = '';

    signatures: string[] = [""];
    premiumSignatures: string[] = [];

    premiumSig: string = "";

    constructor(http: Http) {
        http.get('/autotune/signatures.json')
            .map(res => res.json())
            .subscribe(signatures => {
                this.signatures = this.signatures.concat(signatures.Open).concat(signatures.Startup);
                this.premiumSignatures = this.premiumSignatures.concat(signatures.Startup);
            });
    }

    selected($event, sig) {
        if(this.premiumSignatures.some(x => sig == x)) {
            this.premiumSig = sig + " is not available in the Open version and must be purchased. Please see the pricing for more information."
        } else {
            this.premiumSig = "";
        }
    }
}
