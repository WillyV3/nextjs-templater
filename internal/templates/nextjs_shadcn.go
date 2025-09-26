package template

type Item struct {
	Id          int
	Title       string
	Desc        string
	Command     string
	CommandArgs string
}

var NEXTJS_SHADCN_TEMPLATES = []Item{
	{Id: 0, Title: "nextjs-default", Desc: "Next.js 15 + Shadcn/ui (default theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-default $(pwd)"},
	{Id: 1, Title: "nextjs-modern-minimal", Desc: "Next.js 15 + Shadcn/ui (modern-minimal theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-modern-minimal $(pwd) modern-minimal"},
	{Id: 2, Title: "nextjs-violet-bloom", Desc: "Next.js 15 + Shadcn/ui (violet-bloom theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-violet-bloom $(pwd) violet-bloom"},
	{Id: 3, Title: "nextjs-t3-chat", Desc: "Next.js 15 + Shadcn/ui (t3-chat theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-t3-chat $(pwd) t3-chat"},
	{Id: 4, Title: "nextjs-mocha-mousse", Desc: "Next.js 15 + Shadcn/ui (mocha-mousse theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-mocha-mousse $(pwd) mocha-mousse"},
	{Id: 5, Title: "nextjs-amethyst-haze", Desc: "Next.js 15 + Shadcn/ui (amethyst-haze theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-amethyst-haze $(pwd) amethyst-haze"},
	{Id: 6, Title: "nextjs-doom-64", Desc: "Next.js 15 + Shadcn/ui (doom-64 theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-doom-64 $(pwd) doom-64"},
	{Id: 7, Title: "nextjs-kodama-grove", Desc: "Next.js 15 + Shadcn/ui (kodama-grove theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-kodama-grove $(pwd) kodama-grove"},
	{Id: 8, Title: "nextjs-cosmic-night", Desc: "Next.js 15 + Shadcn/ui (cosmic-night theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-cosmic-night $(pwd) cosmic-night"},
	{Id: 9, Title: "nextjs-quantum-rose", Desc: "Next.js 15 + Shadcn/ui (quantum-rose theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-quantum-rose $(pwd) quantum-rose"},
	{Id: 10, Title: "nextjs-bold-tech", Desc: "Next.js 15 + Shadcn/ui (bold-tech theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-bold-tech $(pwd) bold-tech"},
	{Id: 11, Title: "nextjs-elegant-luxury", Desc: "Next.js 15 + Shadcn/ui (elegant-luxury theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-elegant-luxury $(pwd) elegant-luxury"},
	{Id: 12, Title: "nextjs-amber-minimal", Desc: "Next.js 15 + Shadcn/ui (amber-minimal theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-amber-minimal $(pwd) amber-minimal"},
	{Id: 13, Title: "nextjs-neo-brutalism", Desc: "Next.js 15 + Shadcn/ui (neo-brutalism theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-neo-brutalism $(pwd) neo-brutalism"},
	{Id: 14, Title: "nextjs-solar-dusk", Desc: "Next.js 15 + Shadcn/ui (solar-dusk theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-solar-dusk $(pwd) solar-dusk"},
	{Id: 15, Title: "nextjs-pastel-dreams", Desc: "Next.js 15 + Shadcn/ui (pastel-dreams theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-pastel-dreams $(pwd) pastel-dreams"},
	{Id: 16, Title: "nextjs-clean-slate", Desc: "Next.js 15 + Shadcn/ui (clean-slate theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-clean-slate $(pwd) clean-slate"},
	{Id: 17, Title: "nextjs-ocean-breeze", Desc: "Next.js 15 + Shadcn/ui (ocean-breeze theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-ocean-breeze $(pwd) ocean-breeze"},
	{Id: 18, Title: "nextjs-retro-arcade", Desc: "Next.js 15 + Shadcn/ui (retro-arcade theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-retro-arcade $(pwd) retro-arcade"},
	{Id: 19, Title: "nextjs-midnight-bloom", Desc: "Next.js 15 + Shadcn/ui (midnight-bloom theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-midnight-bloom $(pwd) midnight-bloom"},
	{Id: 20, Title: "nextjs-northern-lights", Desc: "Next.js 15 + Shadcn/ui (northern-lights theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-northern-lights $(pwd) northern-lights"},
	{Id: 21, Title: "nextjs-vintage-paper", Desc: "Next.js 15 + Shadcn/ui (vintage-paper theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-vintage-paper $(pwd) vintage-paper"},
	{Id: 22, Title: "nextjs-sunset-horizon", Desc: "Next.js 15 + Shadcn/ui (sunset-horizon theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-sunset-horizon $(pwd) sunset-horizon"},
	{Id: 23, Title: "nextjs-starry-night", Desc: "Next.js 15 + Shadcn/ui (starry-night theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-starry-night $(pwd) starry-night"},
	{Id: 24, Title: "nextjs-soft-pop", Desc: "Next.js 15 + Shadcn/ui (soft-pop theme)", Command: "bash", CommandArgs: "create-nextjs-shadcn.sh nextjs-soft-pop $(pwd) soft-pop"},
}