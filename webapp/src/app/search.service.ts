import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { Movie } from "./models/movie";

@Injectable({
  providedIn: "root",
})
export class SearchService {
  constructor(private readonly httpClient: HttpClient) {}

  searchMovies(phrase: string): Observable<Movie[]> {
    return this.generalSearch<Movie>(phrase, "movies");
  }
  searchActors(phrase: string): Observable<Movie[]> {
    return this.generalSearch<Movie>(phrase, "actors");
  }
  private generalSearch<T>(
    phrase: string,
    searchtype: string
  ): Observable<T[]> {
    return this.httpClient.get<T[]>(`${environment.apiUrl}/api/search`, {
      params: { phrase, searchtype },
    });
  }
}
