import {Component, OnInit} from '@angular/core';
import {Deposit} from '../../models/deposit';
import {HttpClient, HttpHeaders} from '@angular/common/http';

// const opts: HttpHeaders = {'Content-Type':'application/json'}

@Component({
  selector: 'app-deposit',
  templateUrl: './deposit.component.html',
  styleUrls: ['./deposit.component.css']
})
export class DepositComponent implements OnInit {
  dep: Deposit = <Deposit>{};

  constructor(private http: HttpClient) {
  }

  // TODO: add catching error
  deposit() {
    this.http.post('http://localhost:8080/deposit', JSON.stringify(this.dep)).subscribe((data: any) => {
      console.log(data);
    });
  }

  ngOnInit() {
  }

}
