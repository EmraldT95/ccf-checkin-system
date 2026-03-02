export interface PersonModel {
  id: string
  Name: string
}

export class Person {
  id: string
  Name: string

  constructor(data?: PersonModel) {
    this.id = data?.id ?? ''
    this.Name = data?.Name ?? ''
  }
}
