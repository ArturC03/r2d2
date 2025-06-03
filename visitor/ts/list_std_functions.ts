function extractFunctionsFromObject(obj: any, prefix: string = "", visited = new Set()): Record<string, string[]> {
  const result: Record<string, string[]> = {};

  if (!obj || typeof obj !== "object") return result;
  if (visited.has(obj)) return result;

  visited.add(obj);

  for (const key of Object.getOwnPropertyNames(obj)) {
    try {
      const fullName = prefix ? `${prefix}.${key}` : key;
      const value = obj[key];

      if (typeof value === "function") {
        const argMatch = value.toString().match(/^[^(]*\(([^)]*)\)/);
        const args = argMatch ? argMatch[1].split(',').map(s => s.trim()).filter(Boolean) : [];
        result[fullName] = args;
      } else if (typeof value === "object" && value !== null) {
        Object.assign(result, extractFunctionsFromObject(value, fullName, visited));
      }
    } catch {
      // ignore properties that throw
    }
  }

  return result;
}

const roots = [
  globalThis,
  globalThis.console,
  globalThis.document,
  globalThis.navigator,
  globalThis.window,
  globalThis.performance,
];

const allFunctions: Record<string, string[]> = {};

for (const root of roots) {
  Object.assign(allFunctions, extractFunctionsFromObject(root));
}

console.log(JSON.stringify(allFunctions, null, 2));
