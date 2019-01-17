import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Component({
  selector: 'app-exit',
  templateUrl: './exit.component.html',
  styleUrls: ['./exit.component.css']
})
export class ExitComponent implements OnInit {

  constructor(private http: HttpClient) {
  }

  exit() {
    this.http.post('http://localhost:8080/exit', null).subscribe((data: any) => {
      console.log(data);
    });
  }

  ngOnInit() {
  }

}
