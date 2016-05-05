import {bootstrap}  from 'angular2/platform/browser';
import {HTTP_PROVIDERS, Http} from 'angular2/http';

import {AutotuneCost} from './autotune_cost.ts';
import {AutotuneCurl} from './autotune_curl.ts'


if(document.getElementsByTagName("autotune-cost").length > 0) {
    bootstrap(AutotuneCost);
}

if(document.getElementsByTagName("autotune-curl").length > 0) {
    bootstrap(AutotuneCurl, [
        HTTP_PROVIDERS,
    ]);
}
