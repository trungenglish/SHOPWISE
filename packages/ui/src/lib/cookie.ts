export function getCookie(name: string): string | null {
  if (typeof document === "undefined") {
    return null;
  }

  const prefix = `${name}=`;
  const cookie = document.cookie
    .split(";")
    .map((value) => value.trimStart())
    .filter((value) => value.startsWith(prefix))[0];

  if (!cookie) {
    return null;
  }

  return cookie.substring(prefix.length);
}

const shouldSetCookieDomain = (domain: string): boolean => {
  const d = domain.trim().toLowerCase();
  if (!d) return false;
  // Host-only cookies on loopback: `Domain=localhost` is invalid/rejected in many browsers
  return d !== "localhost" && d !== "127.0.0.1";
};

export function setCookie(
  name: string,
  value: string,
  domain: string,
  expires = 0,
  path = "/"
): void {
  if (typeof document === "undefined") {
    return undefined;
  }

  let cookieString = `${name}=${value}; path=${path}; domain=${domain};`;

  if (expires) {
    const date = new Date();
    date.setTime(date.getTime() + expires * 24 * 60 * 60 * 1000);

    cookieString += ` expires=${date.toUTCString()};`;
  }

  document.cookie = cookieString;

  // Returning undefined because of ESLint's "consistent-return" rule
  return undefined;
}

export function clearCookie(name: string, domain: string, path = "/"): void {
  if (typeof document === "undefined") {
    return undefined;
  }

  let cookieString = `${name}=; path=${path}`;
  if (shouldSetCookieDomain(domain)) {
    cookieString += `; domain=${domain}`;
  }
  cookieString += "; expires=Thu, 01 Jan 1970 00:00:00 GMT";
  document.cookie = cookieString;

  return undefined;
}
