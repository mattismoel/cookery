export const generateRandomTrackingId = () => {
  return crypto.randomUUID() as string;
}
