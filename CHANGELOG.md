# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## (Unreleased) [0.4.0](https://github.com/trento-project/trento/releases/tag/0.4.0) 2021-09-15

## [0.3.0](https://github.com/trento-project/trento/releases/tag/0.3.0) 2021-07-14

### Added

- Check that the HANA and HANA SPS versions are compatible (#100)
- Check NTP time synchronization is configured (#105)
- Get system replication and landscape data for HANA (#106)
- Add self-hosted actions runner continuous deployment (#108)
- Add Check the hacluster user's password is not linux (#109)
- Add version command to CLI (#111)
- Add: Set log timestamp format (#113)
- Generate new consistent ID and human readable name for clusters (#114)
- Collect hdbnsutil -sr_state output (#115)
- Add resources members happy path tests (#119)
- Add: Check that HANA's autostart is disabled (#121)
- Add cloud metadata details - Azure (#126)
- Add pagination to the hosts view table (#132)
- Implement new styles for the host list view (#139)
- List clusters template revamp (#142)

### Changed

- Change: standard logging --> github.com/sirupsen/logrus (#107)
- Update hana-scale-up-perf-optimized-azure.yaml (#110)
- Single cluster template revamp (#120)
- Separate tags in the host list view (#123)
- Update generic layout styles to make the view wider (#138)

### Fixed

- Fix empty tables with a user-friendly message in case of no records (#102)
- Fix minor text polishings (#116)
- Lock mockery dependency version (#121)
- Fix broken test due to cluster-id introduction (#125)
- Add templates sanity check (#129)
- Fix TestSAPSystemHandler test (#130)
- Side menu: show the right title on hover (#131)
- Fix missing `</li>` in menu sidebar (#133)

## [0.2.0](https://github.com/trento-project/trento/releases/tag/0.2.0) 2021-06-16

### Added

- Add SAP Systems to default environment and landscape in absence of one (#70)
- Check that /etc/hosts contains all cluster nodes (#98)
- Check the UCAST is used by corosync with at least 2 com-n rings (#91, #96)
- Add project logo to the header (#90)
- Check that 2-nodes cluster must either have disk-based sbd or qdevice (#87)
- Landing page update with scope documentation (#82)
- Add this changelog ;) (#80)
- SBD configuration and service discovery (#72)

### Changed

- README updates
- Side bar and Home landing page improvements
- azure-rules check 1.3.5 was splitted into two checks
- Improve sidebar template (#84)
- Copy sapcontrol webservices from the exporter library instead of importing them (#81)
- Change how some checks are grouped together (#73, #74, #94)
- Reorganize the SAP System domain model structures (#75)

### Fixed

- Fix SAP system layout rendering
- Don't let the app crash on 404s (#97)
- Use the correct path for the global.ini config file of the SAPHanaSR check (#95)

## [0.1.0](https://github.com/trento-project/trento/releases/tag/0.1.0) 2021-05-26

### Added

- first release of Trento
- Automated discovery of SAP HANA HA clusters
- SAP Systems and Instances overview
- Grouping by Landscapes and Environments
- Configuration validation for Pacemaker, Corosync, SBD, SAPHanaSR and other + generic SUSE Linux Enterprise for SAP Application OS settings
- Specific configuration audits for SAP HANA Scale-Up Performance-Optimized
- scenarios deployed on MS Azure cloud.
