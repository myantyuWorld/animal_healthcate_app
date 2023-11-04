import { AxiosPromise } from "axios";
import repository from "./repository";

const resource = "/pet";

export interface Pet {
  name: string;
  age: string;
  sex: string;
  nowWeight: string;
  targetWeight: string;
  birthDay: string;
}

export default class PetRepository {
  // TODO: AxiosPromise<Pet>にならない？
  // Express+TypeScriptでレスポンスボディに「ストレスなく」型を付ける | https://techblog.istyle.co.jp/archives/8568
  // [TypeScript] Axiosのtry/catchでの例外オブジェクトを型付けする | https://dev.classmethod.jp/articles/typescript-typing-exception-objects-in-axios-trycatch/
  public getPetSummary(id: number): AxiosPromise<Pet> {
    return repository.get<Pet>(`${resource}/${id}`);
  }
}
