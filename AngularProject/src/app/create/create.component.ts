import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { Http } from '@angular/http';

@Component({
    selector: 'app-create',
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {

    public timer: any;

    public constructor(private location: Location, private http: Http) {
        this.timer = {
            "name": "",
            "data": {
                "sets": false,
                "work": false,
                "rest": false
            }
        }
    }

    public ngOnInit() { }

    public save() {
        if(this.timer.name) {
            this.http.post("http://localhost:12345/timer", JSON.stringify(this.timer))
                .subscribe(result => {
                    this.location.back();
                });
        }
    }

}
