import { TimerComponent } from "./timer/timer.component";
import { CreateComponent } from "./create/create.component";
import { LoadComponent } from "./load/load.component";

export const AvailableRoutes: any = [
  { path: "", component: TimerComponent },
  { path: "create", component: CreateComponent },
  { path: "load", component: LoadComponent }
];
