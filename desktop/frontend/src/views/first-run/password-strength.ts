export type strength = "very-weak" | "weak" | "strong" | "very-strong";


export function getPasswordStrength(password: string) : strength {
  let score = 0;

  if (!password) return "very-weak";
  if (password.length >= 8) score++;
  if (password.length >= 12) score++;

  if (/[a-z]/.test(password)) score++;
  if (/[A-Z]/.test(password)) score++;
  if (/\d/.test(password)) score++;
  if (/[\W_]/.test(password)) score++;

  if (score <= 2) return "very-weak";
  if (score <= 4) return "weak";
  if (score <= 5) return "strong";
  return "very-strong";
}
