export type Unit = "kg" | "g" | "L" | "mL" | "dL" | "cL";

export type Ingredient = {
  name: string;
  amount?: number;
  unit?: Unit;
  note?: string;
}

export type SubRecipe = {
  title?: string;
  ingredients: Ingredient[];
  instructions: Instruction[];
}

export type Recipe = {
  title: string;
  bannerUrl: string;
  author: string;
  cookMinutes: number;
  totalMinutes: number;
  subRecipes: SubRecipe[];
}

export type Instruction = {
  text: string;
  imageUrl?: string;
}
