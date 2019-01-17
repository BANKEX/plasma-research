import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Transfer} from '../../models/transfer';

@Component({
  selector: 'app-transfer',
  templateUrl: './transfer.component.html',
  styleUrls: ['./transfer.component.css']
})
export class TransferComponent implements OnInit {
  tx: Transfer = <Transfer>{};

  constructor(private http: HttpClient) {
  }

  transfer() {
    this.http.post('http://localhost:8080/transfer', JSON.stringify(this.tx)).subscribe((data: any) => {
      console.log(data);
    });
  }

  ngOnInit() {
  }

}
