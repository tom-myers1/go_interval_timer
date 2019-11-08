import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';
import { Router } from '@angular/router'
import { Location } from '@angular/common'

@Component({
  selector: 'app-timer',
  templateUrl: './timer.component.html',
  styleUrls: ['./timer.component.css']
})
export class TimerComponent implements OnInit {

  public timer: any;

  public constructor(private http: Http, private router: Router, private location: location) {
    this.timer = [];
  }

  public ngOnInit() {
    this.location.subscribe(() => {
      this.refresh();
    });
    this.refresh();
  }

  private refresh() {
    this.http.get("http://localhost:12345/timer")
      .map(result => result.json())
      .subscribe(result => {
        this.timer = result;
      });
  }

  public search(event: any) {
    let url = "http://localhost:12345/timer";
    if (event.target.value) {
      url = "http://localhost:12345/search/" + event.target.value;
    }
    this.http.get(url)
      .map(result => result.json())
      .subscribe(result => {
        this.timer = result;
      });
  }

  public create() {
    this.router.navigate(["create"]);
  }

  public load() {
    this.router.navigate(["load"]);
  }
}
