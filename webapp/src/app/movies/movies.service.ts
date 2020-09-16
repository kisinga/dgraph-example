import { Injectable } from "@angular/core";
import { Movie } from "../models/movie";
import { HttpClient } from "@angular/common/http";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class MoviesService {
  constructor(private readonly httpClient: HttpClient) {}

  getMovies(): Observable<Movie[]> {
    return this.httpClient.get<Movie[]>(`${environment.apiUrl}/api/users`);
  }
}
