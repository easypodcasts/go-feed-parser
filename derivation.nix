{ lib, buildGoPackage}:
let
  yorodm = {
    name = "Yoandy Rodriguez" ;
    email = "no_spam@nospam.org";
    github = "yorodm";
  };
in
buildGoPackage rec {
  name = "go-feed-parser-${version}";
  version = "v.0.0.1";
  goPackagePath = "github.com/easypodcasts/go-feed-parser";

  src = lib.cleanSource ./.;

  meta = {
    description = "Simple feed parser for Easypodcasts";
    homepage = https://github.com/nix-community/queued-build-hook;
    maintainers = [ yorodm ];
  };

}
