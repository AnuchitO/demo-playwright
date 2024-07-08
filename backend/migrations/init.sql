CREATE TABLE IF NOT EXISTS skill (
	key TEXT PRIMARY KEY,
	name TEXT NOT NULL DEFAULT '',
	description TEXT NOT NULL DEFAULT '',
	logo TEXT NOT NULL DEFAULT '',
	levels JSONB NOT NULL DEFAULT '[]',
	tags TEXT [] NOT NULL DEFAULT '{}'
);

INSERT INTO
	skill (key, name, description, levels, tags)
VALUES
	(
		'go',
		'Go',
		'Go is an open source programming...',
		'[{"name": "Beginner", "level": 1, "descriptions": ["basic knowledge ..."]},{"name": "Intermediate", "level": 2, "descriptions": ["complex programs..."]}]',
		'{go,golang}'
	) ON CONFLICT (key) DO UPDATE
SET
	name = EXCLUDED.name,
	description = EXCLUDED.description,
	levels = EXCLUDED.levels,
	tags = EXCLUDED.tags;

INSERT INTO
	skill (key, name, description)
VALUES
	(
		'figma',
		'Figma',
		'Figma is a vector graphics editor and prototyping tool which is primarily web-based, with additional offline features enabled by desktop applications for macOS and Windows.'
	),
	(
		'html5',
		'HTML5',
		'HTML5 is a markup language used for structuring and presenting content on the World Wide Web.'
	),
	(
		'negotiation',
		'Negotiation',
		'Negotiation is a dialogue between two or more people or parties intended to reach a beneficial outcome over one or more issues where a conflict exists with respect to at least one of these issues.'
	) ON CONFLICT (key) DO NOTHING;
