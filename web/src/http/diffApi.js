import client from "./client";

export async function getV1DiffTarget(originURL, compareURL, method, bodies, headers) {
  const encodeOrigin = encodeURIComponent(originURL)
  const encodeCompare = encodeURIComponent(compareURL)
  const bodyJSON = encodeURIComponent(JSON.stringify(bodies))
  return await client().get(`/v1/diff-target?originURL=${encodeOrigin}&compareURL=${encodeCompare}&method=${method}&bodyJSON=${bodyJSON}&headerKeys=${Object.keys(headers).join(',')}`, {
    headers: {
      ...headers,
    }
  })
}
