type User {
    required property email -> str {
        constraint exclusive;
    };
    property password -> str;
    property dob -> str;
    property verified -> bool;
    property firstname -> std::optional<std::str>;  // Mark firstname as optional
    property lastname -> str;
    property phone -> str;
    property cv -> str;
    property accesstoken -> str;
    property refreshtoken -> str;
}
