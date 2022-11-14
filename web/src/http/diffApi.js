import client from "./client";

export async function getV1DiffTarget(originURL, compareURL, headers) {
  return await client().get(`/v1/diff-target?originURL=${encodeURIComponent(originURL)}&compareURL=${encodeURIComponent(compareURL)}&headerKeys=${Object.keys(headers).join(',')}`, {
    headers: {
      ...headers,
    }
  })
}
