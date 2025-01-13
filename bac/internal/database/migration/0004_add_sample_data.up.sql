INSERT INTO resources (
    id,
    name,
    description,
    latitude,
    longitude,
    diagnoses,
    address,
    contact_info
) VALUES 
(
    gen_random_uuid(),  -- Let PostgreSQL generate the UUID
    'Mental Health Center LA',
    'Comprehensive mental health services for adults and children',
    34.0522,
    -118.2437,
    ARRAY['ADHD', 'Anxiety', 'Depression'],
    '123 Main St, Los Angeles, CA 90012',
    '{"phone": "(555) 123-4567", "email": "info@mhcla.org", "hours": "Mon-Fri 9am-5pm"}'::jsonb
),
(
    gen_random_uuid(),
    'Autism Support Center',
    'Specialized autism support services and therapy',
    34.0511,
    -118.2428,
    ARRAY['Autism'],
    '456 Oak St, Los Angeles, CA 90013',
    '{"phone": "(555) 987-6543", "email": "support@autismcenter.org", "hours": "Mon-Sat 8am-6pm"}'::jsonb
),
(
    gen_random_uuid(),  -- Let PostgreSQL generate the UUID
    'Integrated Care Clinic',
    'Multi-disciplinary mental health care',
    34.0458,
    -118.2661,
    ARRAY['ADHD', 'Depression', 'Anxiety', 'Autism'],
    '789 Elm St, Los Angeles, CA 90014',
    '{"phone": "(555) 246-8135", "email": "care@icc.org", "hours": "24/7", "emergencyLine": "(555) 999-9999"}'::jsonb
);
