import { inject } from "vue";
import { formContextKey } from "../context/formcontext";

export const useFormContext = () => {
  const context = inject(formContextKey) as any;
  if (!context) {
    throw new Error(
      "Please Call The Controller Component Component Under RxForm"
    );
  }
  return context;
};
