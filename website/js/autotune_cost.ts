import 'rxjs/Rx';
import {Component} from 'angular2/core';
import {NgForm}    from 'angular2/common';

@Component({
    selector: 'autotune-cost',
    // Location of the template for this component
    template: `
<div id="barchart_values"></div>

<div>
  <label>Number of Machines:</label>
  <input type="range" min="1" max="1000" step="1" [(ngModel)]="numMachines"  />
</div>

<caption>Cost for m4.large</caption>
<table class="table">
<tr>
<th></th>
<th>OnDemand Pricing<th>
</tr>
<tr>
<td>Number of Machines</td>
<td>{{numMachines}}</td>
</tr>
<tr>
<td>AWS m4.large</td>
<td>\${{machinesPricing}} per month</td>
</tr>
<tr>
<td>Savings</td>
<td>\${{savings}} per month</td>
</tr>
</table>
`

})
export class AutotuneCost {
    selectedProfile: string = 'apache';


    constructor() {
        google.charts.load("current", {packages:["corechart"]});
        google.charts.setOnLoadCallback(this.drawChart);

        this.numMachines = 10;
    }

    drawChart() {
        let data = google.visualization.arrayToDataTable([
            ["Tuning State", "Machines", { role: "style" } ],
            ["Pre-Tuning", 100, "red"],
            ["Post-Tuning 5%", 95, "green"],
            ["Post-Tuning 15%", 85, "green"],
            ["Post-Tuning 25%", 75, "green"],
        ]);

        let view = new google.visualization.DataView(data);
        view.setColumns([
            0,
            1,
            { calc: "stringify",
              sourceColumn: 1,
              type: "string",
              role: "annotation" },
            2]);

        let options = {
            title: "Number of Machines Saved",
            width: 600,
            height: 400,
            legend: { position: "none" },
        };
        let chart = new google.visualization.BarChart(document.getElementById("barchart_values"));
        chart.draw(view, options);
    }

    get numMachines() {
        return this.numMachinesInt;
    }

    set numMachines(value) {
        this.numMachinesInt = value

        this.m4LargePricing = 60.59
        this.savedPercent = 0.25;
        this.machinesPricing = this.m4LargePricing * this.numMachines;
        this.savings = this.numMachines * this.m4LargePricing - (this.numMachines * (1 - this.savedPercent)) * this.m4LargePricing;
    }
}
