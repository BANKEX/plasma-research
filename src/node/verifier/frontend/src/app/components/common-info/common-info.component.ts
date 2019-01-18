import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {CommonInfo} from '../../models/commonInfo';
import {Observable} from 'rxjs';
import {timer} from 'rxjs';
import {take} from 'rxjs/operators';

@Component({
  selector: 'app-common-info',
  templateUrl: './common-info.component.html',
  styleUrls: ['./common-info.component.css']
})
export class CommonInfoComponent implements OnInit {
  cI: CommonInfo = <CommonInfo>{};

  constructor(private http: HttpClient) {
  }

  ngOnInit() {
    timer(5000, 1000).pipe(
      take(1000)).subscribe(x => {
      this.getInfo();
    });
  }

  getInfo() {
    this.http.get('http://localhost:8080/common').subscribe((data: any) => {
      this.cI.contractAddress = data.contract_address;
      this.cI.contractBalance = data.contract_balance;
      this.cI.verifierPlasmaBalance = data.verifier_plasma_balance;
      this.cI.verifierEtherBalance = data.verifier_ether_balance;
      this.cI.latestBlock = data.latest_block;
      this.cI.verifierInputs = data.verifier_inputs;
    });
  }

}
