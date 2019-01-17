import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {History} from '../../models/history';
import {MatTableDataSource} from '@angular/material';
import {timer} from 'rxjs';
import {take} from 'rxjs/operators';

@Component({
  selector: 'app-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.css']
})
export class HistoryComponent implements OnInit {
  h: History[] = [];
  dataSource = new MatTableDataSource(this.h);

  constructor(private http: HttpClient) {
  }

  ngOnInit() {
    this.h = [];
    this.getInfo();
    timer(10000, 1000).pipe(
      take(1000)).subscribe(x => {
      this.h = [];
      this.getInfo();
    });
  }

  getInfo() {
    this.http.get('http://localhost:8080/operations').subscribe((data: any) => {
      for (let i = 0; i < data.Events.length; i++) {
        this.h.push(data.Events[i]);
      }
      console.log(this.h);
    });
  }

}
