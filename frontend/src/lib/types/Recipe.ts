
export type Recipe = {
  id?: number;
  title: string;
  bannerUrl: string;
  author: string;
  cookMinutes: number;
  totalMinutes: number;
  subRecipes: SubRecipe[];
}


export type Ingredient = {
  id?: number;
  name: string;
  amount?: number;
  unit?: string;
  note?: string;
}

export type SubRecipe = {
  id?: number;
  title?: string;
  ingredients: Ingredient[];
  instructions: Instruction[];
}


export type Instruction = {
  id?: number;
  text: string;
  imageUrl?: string;
}

export type TrackableSubRecipe = SubRecipe & { trackingId: string };
export type TrackableIngredient = Ingredient & { trackingId: string }
export type TrackableInstruction = Instruction & { trackingId: string }
