import { HandlerContext } from "$fresh/server.ts";

interface Release {
  tag_name: string;
}

export const handler = async (_req: Request, _ctx: HandlerContext) => {
  const { user, repo } = _ctx.params;

  const u = new URL(_req.url);

  const resp = await fetch(
    `https://api.github.com/repos/${user}/${repo}/releases/latest`,
    {
      headers: {
        Authorization: u.searchParams.get("token")
          ? `token ${u.searchParams.get("token")}`
          : "",
      },
    }
  );

  if (resp.status === 404) {
    return new Response(null);
  }

  const release: Release = await resp.json();

  let tag = release.tag_name;

  if (u.searchParams.get("no-v")) {
    tag = tag.replace("v", "");
  }

  return new Response(tag);
};
