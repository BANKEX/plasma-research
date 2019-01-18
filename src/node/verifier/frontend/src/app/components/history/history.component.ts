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
    timer(10000, 10000).pipe(
      take(1000)).subscribe(x => {
      this.h = [];
      this.getInfo();
    });
  }

  getInfo() {
    this.http.get('http://localhost:8080/history').subscribe((data: any) => {
      for (let i = 0; i < data.Events.length; i++) {
        this.h.push(data.Events[i]);
      }
    });
  }

  date(timestamp: number): string {
    timestamp = timestamp * 1000;
    const date: number = new Date(timestamp).getDate();
    const month: number = new Date(timestamp).getMonth() + 1;
    const year: number = new Date(timestamp).getFullYear();
    return month + '/' + date + '/' + year;
  }

}
